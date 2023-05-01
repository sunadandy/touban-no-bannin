<template>
  <div>
  </div>
</template>

<script>
import { parse, format, startOfToday, startOfWeek, add } from 'date-fns'

export default {
  name: 'CalcMemberInfo',
  methods:{
    CalcuNewMemberInfo(toubanInfo, selected){
      var new_memberInfos = []
      var currentMaxOrder = 0
      for(var i=0; i<selected.length; i++){
        var memberInfo = {}
        memberInfo.touban_id = toubanInfo.id
        memberInfo.affiliation = selected[i].affiliation
        memberInfo.name = selected[i].name
        memberInfo.emplyoee_number = "E" + selected[i].emplyoeeNo   //バックエンド側で数値文字列を上手く扱えないのでE付きに。
        memberInfo.order_number = currentMaxOrder + 1
        memberInfo.last = "2099-12-31"  //適当な初期値
        memberInfo.next = null          //後で計算
        memberInfo.email = selected[i].email
        new_memberInfos.push(memberInfo)

        currentMaxOrder = memberInfo.order_number
      }
      // メンバー分の次回実施日を計算
      const data = toubanInfo.scheduling.split("-")
      const interval = parseInt(data[0])
      const day = parseInt(data[1])
      const startDate = toubanInfo.start
      new_memberInfos = this.ReSchedule(new_memberInfos, interval, day, startDate)

      return new_memberInfos
    },
    // メンバーの入れ替え（※入れ替えのみ行う）
    ReplaceMember(current_memberInfos, selected){
      var memberInfos = []

      const toubanId = parseInt(current_memberInfos[0].touban_id)   //要素番号はなんでもいい
      // 現在のorderの最大値を取得
      var currentMaxOrder = current_memberInfos.reduce((prev, current) => (prev.order_number > current.order_number) ? prev : current).order_number;

      // 既存メンバーと新しく選択されたメンバーを照合
      for(var j=0; j<selected.length; j++){
        for(var i=0; i<current_memberInfos.length; i++){
          // 一致する場合は既存メンバーなので既存データをそのままpush
          if(current_memberInfos[i].name == selected[j].name){
            memberInfos.push(current_memberInfos[i])
            break
          }
          // 最後まで一致する名前が見つからなかった場合は新規メンバー
          if(i == current_memberInfos.length-1){
            var memberInfo = {}
            memberInfo.touban_id = toubanId
            memberInfo.affiliation = selected[j].affiliation
            memberInfo.name = selected[j].name
            memberInfo.emplyoee_number = "E" + selected[j].emplyoeeNo   //バックエンド側で数値文字列を上手く扱えないのでE付きに。
            memberInfo.order_number = currentMaxOrder + 1
            memberInfo.last = "2099-12-31"  //適当な初期値
            memberInfo.next = null          //後で計算
            memberInfo.mail = selected[j].email
            memberInfos.push(memberInfo)

            currentMaxOrder = memberInfo.order_number
          }
        }
      }
      // order_numberでソートして採番
      memberInfos.sort((a, b) => a.order_number - b.order_number);
      for (var i = 0; i < memberInfos.length; i++) {
        memberInfos[i].order_number = i + 1;
      }

      return memberInfos
    },
    ReSchedule(memberInfos, interval, day, startDate){
      const today = format(startOfToday(), 'yyyy-MM-dd') 
      // 今日の日付に一番近い最終実施日のインデックスを検索する
      const closestLastDateIndex = memberInfos.reduce((closestIndex, memberInfo, index) => {
        const parsedCurrentLastDate = parse(memberInfo.last, 'yyyy-MM-dd', new Date());
        const parsedToday  = parse(today, 'yyyy-MM-dd', new Date());
        // 初めてtoday>LastDateが成立するまではこれを繰り返す
        if(closestIndex == index){
          if (parsedToday > parsedCurrentLastDate){
            return index
          }else{
            return index + 1
          }
        }

        const parsedClosestLastDate = parse(memberInfos[closestIndex].last, 'yyyy-MM-dd', new Date());          
        if (parsedToday > parsedCurrentLastDate && parsedCurrentLastDate > parsedClosestLastDate) {
          return index;
        }else{
          return closestIndex;
        }
      }, 0)
      
      // closestLastDateIndexがメンバーJSON配列の長さと同値の場合は新規当番作成、もしくは一度も実施していない
      var next = startDate
      if(closestLastDateIndex == memberInfos.length){ 
        // 不定期以外は、order=1以外のメンバーの予定も設定
        if(interval != 5){
          memberInfos.forEach(memberInfo => {
            memberInfo.next = next
            next = format(startOfWeek(add(parse(next, 'yyyy-MM-dd', new Date()), { weeks: interval }), { weekStartsOn: day }), 'yyyy-MM-dd')
          });
        }
      }else{
        // memberInfos[closestLastDateIndex].orderの次のorderから順次リスケ
        // 最初はmemberInfos[closestLastDateIndex].order < memberInfo.orderをリスケ
        const lastOrder = memberInfos[closestLastDateIndex].order_number
        var nextOrder = lastOrder + 1
        memberInfos.forEach((memberInfo) => {
          if(memberInfo.order_number == nextOrder){
            memberInfo.next = next
            next = format(startOfWeek(add(parse(next, 'yyyy-MM-dd', new Date()), { weeks: interval }), { weekStartsOn: day }), 'yyyy-MM-dd')
            nextOrder++
          }
        })
        // 次にはmemberInfos[closestLastDateIndex].order => memberInfo.orderをリスケ
        var nextOrder = 1
        memberInfos.forEach((memberInfo) => {
          if(memberInfo.order_number <= lastOrder){
            memberInfo.next = next
            next = format(startOfWeek(add(parse(next, 'yyyy-MM-dd', new Date()), { weeks: interval }), { weekStartsOn: day }), 'yyyy-MM-dd')
            nextOrder++
          }
        })
      }
      console.log(memberInfos)
      return memberInfos
    },
  },
}
</script>