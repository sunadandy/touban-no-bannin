<template>
  <div>
  </div>
</template>

<script>
import { addDays, format, parseISO, startOfToday } from 'date-fns'

export default {
  name: 'CalcMemberInfo',
  methods:{
    CalcuNewMemberInfo(toubanId, interval_type, selected){
      var new_memberInfos = []

      // 当番IDに一致する順番レコードを取得
      const current_memberInfos = this.$store.getters.GetMemberByToubanId(toubanId)
      // CreateかUpdateかを取得したレコードの長さで判定
      if(current_memberInfos.length == 0){
        var currentMaxOrder = 0
        var currentMaxNextDate = format(startOfToday(), 'yyyy-MM-dd')
        for(var i=0; i<selected.length; i++){
          var memberInfo = {}
          memberInfo.touban_id = parseInt(toubanId)
          memberInfo.name = selected[i].name
          memberInfo.emplyoee_number = "E" + selected[i].emplyoeeNo   //バックエンド側で数値文字列を上手く扱えないのでE付きに。
          memberInfo.order_number = currentMaxOrder + 1
          memberInfo.last = "2100-1-1"  //適当な初期値
          memberInfo.next = this.CalcNextDate(currentMaxNextDate, parseInt(interval_type))
          memberInfo.email = selected[i].email
          new_memberInfos.push(memberInfo)

          currentMaxOrder = memberInfo.order_number
          currentMaxNextDate = memberInfo.next
        }
      }else{
        // 現在のorderの最大値を取得
        var currentMaxOrder = current_memberInfos.reduce((prev, current) => (prev.order_number > current.order_number) ? prev : current).order_number;
        // 現在の次回実施日の最大値を取得
        var maxEpochTime = Math.max(...current_memberInfos.map(x => new Date(x.next).getTime()));
        var currentMaxNextDate = format(maxEpochTime, 'yyyy-MM-dd')

        // 既存メンバーと新しく選択されたメンバーを照合
        for(var j=0; j<selected.length; j++){
          for(var i=0; i<current_memberInfos.length; i++){
            // 一致する場合は既存メンバーなので既存データをそのままpush
            if(current_memberInfos[i].name == selected[j].name){
              new_memberInfos.push(current_memberInfos[i])
              break
            }
            // 最後まで一致する名前が見つからなかった場合は新規メンバー
            if(i == current_memberInfos.length-1){
              var memberInfo = {}
              memberInfo.touban_id = parseInt(toubanId)
              memberInfo.name = selected[j].name
              memberInfo.emplyoee_number = "E" + selected[j].emplyoeeNo   //バックエンド側で数値文字列を上手く扱えないのでE付きに。
              memberInfo.order_number = currentMaxOrder + 1
              memberInfo.last = "2100-1-1"  //適当な初期値
              memberInfo.next = this.CalcNextDate(currentMaxNextDate, parseInt(interval_type))
              memberInfo.mail = selected[j].email
              new_memberInfos.push(memberInfo)
  
              currentMaxOrder = memberInfo.order_number
              currentMaxNextDate = memberInfo.next
            }
          }
        }
        // order_numberでソートして採番
        new_memberInfos.sort((a, b) => a.order_number - b.order_number);
        for (var i = 0; i < new_memberInfos.length; i++) {
          new_memberInfos[i].order_number = i + 1;
        }
      }
      return new_memberInfos
    },
    CalcNextDate(currentMaxNextDate, interval_type){
      const date = parseISO(currentMaxNextDate)
      switch(interval_type){
        // 毎日
        case 1:
          var nextDate = format(addDays(date, 1), 'yyyy-MM-dd')
          return nextDate
          break
        // 毎週
        case 2:
          var nextDate = format(addDays(date, 7), 'yyyy-MM-dd')
          return nextDate
          break
        // 毎月
        case 3:
          var nextDate = format(addDays(date, 30), 'yyyy-MM-dd')
          return nextDate
          break
        // 未定義
        case 4:
          return "2100-01-01"
          break
      }
    },
  },
}
</script>