<template>
  <div>
  </div>
</template>

<script>
import { parse, format, startOfToday, startOfWeek, startOfMonth, endOfMonth, add, addDays, addMonths } from 'date-fns'

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
      const newInterval = parseInt(data[0])
      const newWeek = parseInt(data[1])
      const newDay = parseInt(data[2])
      const newDate = parseInt(data[3])
      const nextDate = toubanInfo.next
      new_memberInfos = this.ReSchedule(new_memberInfos, newInterval, newWeek, newDay, newDate, nextDate)

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
    ReSchedule(memberInfos, newInterval, newWeek, newDay, newDate, nextDate){
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
      
      // closestLastDateIndexがメンバーJSON配列の長さと同値の場合は新規当番作成、もしくは一度も実施していないので全員の次回実施日を設定する
      var next = nextDate
      if(closestLastDateIndex == memberInfos.length){ 
        // 平日毎日
        if(newInterval == 0){
          memberInfos.forEach(memberInfo => {
            memberInfo.next = next
            var date = parse(next, 'yyyy-MM-dd', new Date())
            // 金曜日かどうか判定
            if(date.getDay()  == 5){
              next = format(addDays(date, 3), 'yyyy-MM-dd')
            }else{
              next = format(addDays(date, 1), 'yyyy-MM-dd')
            }
          });
        }else if(newInterval >= 1 && newInterval <= 3){ //毎週、隔週、3週毎
          memberInfos.forEach(memberInfo => {
            memberInfo.next = next
            next = format(startOfWeek(add(parse(next, 'yyyy-MM-dd', new Date()), { weeks: newInterval }), { weekStartsOn: newDay }), 'yyyy-MM-dd')
          });
        }else if(newInterval == 4){ //毎月
          // 曜日指定
          if(newWeek != 0){
            memberInfos.forEach(memberInfo => {
              memberInfo.next = next
              // nextをdateオブジェクトに変更
              var date = parse(next, 'yyyy-MM-dd', new Date())
              // nextで指定された月の最初の曜日を取得(0から6の整数で、日曜日から土曜日までの順)
              var nextMonth = addMonths(date, 1)
              var firstDayOfNextMonth = startOfMonth(nextMonth).getDay()
              // nextで指定された月の最後の曜日を取得(0から6の整数で、日曜日から土曜日までの順)
              var endDayOfNextMonth = endOfMonth(nextMonth).getDay()
              // newDayが月初めの曜日よりも前の場合は第2週の曜日を設定
              if(newWeek == 1 && firstDayOfNextMonth > newDay){
                next = nextMonth.setDate(7 - firstDayOfNextMonth + newDay + 1)
              }else if(newWeek == 1 && firstDayOfNextMonth <= newDay && 6 >= newDay){   //第1週の成立条件
                next = nextMonth.setDate(newDay - firstDayOfNextMonth + 1)  //日曜が0だとすると月初めが日曜日の場合を考えると+1になる。
              }else if(newWeek == 5 && endDayOfNextMonth < newDay){                         //newDayが月終わりの曜日よりも後の場合は前週の曜日を設定
                next = nextMonth.setDate(endOfMonth(nextMonth).getDate() - 7 + (newDay-endDayOfNextMonth))
              }else{                                                                        //それ以外                
                if(firstDayOfNextMonth > newDay){
                  next = nextMonth.setDate((newWeek-1)*7 + (newDay - firstDayOfNextMonth) + 1 + 7)
                }else{
                  next = nextMonth.setDate((newWeek-1)*7 + (newDay - firstDayOfNextMonth) + 1)
                }
              }
              // フォーマット変換
              next = format(next, 'yyyy-MM-dd')
              console.log(next)
            });
          }else{  //日付指定
            // [todo]newDateの更新は1回でいいはずなので若干冗長
            memberInfos.forEach(memberInfo => {
              memberInfo.next = next
              // nextをdateオブジェクトに変更
              var date = parse(next, 'yyyy-MM-dd', new Date())
              // 日付をnewDateに上書き
              date.setDate(newDate)
              // 毎月決まった日を当番に設定する。土日祝日は考慮されない。
              next = format(addMonths(date, 1), 'yyyy-MM-dd')
            });
          }
        }
      }else{
        // memberInfos[closestLastDateIndex].orderの次のorderから順次リスケ
        // 最初はmemberInfos[closestLastDateIndex].order < memberInfo.orderをリスケ
        const lastOrder = memberInfos[closestLastDateIndex].order_number
        var nextOrder = lastOrder + 1
        memberInfos.forEach((memberInfo) => {
          if(memberInfo.order_number == nextOrder){
            memberInfo.next = next
            if(newInterval == 0){
              var date = parse(next, 'yyyy-MM-dd', new Date())
              // 金曜日かどうか判定
              if(date.getDay()  == 5){
                next = format(addDays(date, 3), 'yyyy-MM-dd')
              }else{
                next = format(addDays(date, 1), 'yyyy-MM-dd')
              }
            }else if(newInterval >= 1 && newInterval <= 4){
              next = format(startOfWeek(add(parse(next, 'yyyy-MM-dd', new Date()), { weeks: newInterval }), { weekStartsOn: newDay }), 'yyyy-MM-dd')
            }
            nextOrder++
          }
        })
        // 次にはmemberInfos[closestLastDateIndex].order => memberInfo.orderをリスケ
        var nextOrder = 1
        memberInfos.forEach((memberInfo) => {
          if(memberInfo.order_number <= lastOrder){
            memberInfo.next = next
            if(newInterval == 0){
              var date = parse(next, 'yyyy-MM-dd', new Date())
              // 金曜日かどうか判定
              if(date.getDay()  == 5){
                next = format(addDays(date, 3), 'yyyy-MM-dd')
              }else{
                next = format(addDays(date, 1), 'yyyy-MM-dd')
              }
            }else if(newInterval >= 1 && newInterval <= 4){
              next = format(startOfWeek(add(parse(next, 'yyyy-MM-dd', new Date()), { weeks: newInterval }), { weekStartsOn: newDay }), 'yyyy-MM-dd')
            }
            nextOrder++
          }
        })
      }
      return memberInfos
    },
  },
}
</script>