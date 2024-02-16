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

      function getNextDate(currentDate, increment) {
        const date = parse(currentDate, 'yyyy-MM-dd', new Date());
        return format(addDays(date, increment), 'yyyy-MM-dd');
      }

      function getNextMonthDate(currentDate, newWeek, newDay) {
        const date = parse(currentDate, 'yyyy-MM-dd', new Date());
        // nextで指定された月の最初の曜日を取得(0から6の整数で、日曜日から土曜日までの順)
        const nextMonth = addMonths(date, 1);
        const firstDayOfNextMonth = startOfMonth(nextMonth).getDay();
        // nextで指定された月の最後の曜日を取得(0から6の整数で、日曜日から土曜日までの順)
        const endDayOfNextMonth = endOfMonth(nextMonth).getDay();
        
        if(newWeek == 1){
          // newDayが月初めの曜日よりも前の場合は7日オフセット
          return format(nextMonth.setDate((firstDayOfNextMonth > newDay ? 7 : 0) + (newDay - firstDayOfNextMonth) + 1), 'yyyy-MM-dd');
        } else if (newWeek == 5) {
          //newDayが月終わりの曜日よりも後の場合は前週の曜日を設定
          return format(nextMonth.setDate(endOfMonth(nextMonth).getDate() - (endDayOfNextMonth < newDay ? 7 : 0) - (endDayOfNextMonth - newDay)), 'yyyy-MM-dd');
        } else {
          return format(nextMonth.setDate((newWeek * 7 - 6) + (newDay - firstDayOfNextMonth)), 'yyyy-MM-dd');
        }
      }

      // 今日の日付に一番近い最終実施日のインデックスを検索する（順番入れ替えを考慮するとlastは必ずしも昇順ではない。）
      // reduceは、コールバック関数と初期値を引数に持つ。初期値はコールバック関数の第1引数
      // コールバック関数が最後に return した値が、 reduce() 全体の戻り値になる。
      const closestLastDateIndex = memberInfos.reduce((closestIndex, memberInfo, index) => {
        const parsedCurrentLastDate = parse(memberInfo.last, 'yyyy-MM-dd', new Date());
        const parsedToday  = parse(today, 'yyyy-MM-dd', new Date());
        // 0番目のlastが今日よりも前ならclosestIndexとしてリターン。そうでないならインデックスを進める。このとき、次回もclosestIndex=indexとなるのでこのパスを通る。
        if(closestIndex == index){
          return parsedToday > parsedCurrentLastDate ? index : index+1 
        }
        // memberInfos[index].last と parsedClosestLastDateを比較。parsedClosestLastDateの方がtodayに近い場合はclosestIndexを更新
        const parsedClosestLastDate = parse(memberInfos[closestIndex].last, 'yyyy-MM-dd', new Date());          
        if (parsedToday > parsedCurrentLastDate && parsedCurrentLastDate > parsedClosestLastDate) {
          return index;
        }else{
          return closestIndex;
        }
      }, 0)
      
      var next = nextDate
      // closestLastDateIndexがメンバーJSON配列の長さと同値の場合は新規当番作成、もしくは一度も実施していないので全員の次回実施日を設定する
      if(closestLastDateIndex == memberInfos.length){
        memberInfos.forEach(memberInfo => {
          memberInfo.next = next
          // 平日毎日
          if(newInterval == 0){
            const date = parse(next, 'yyyy-MM-dd', new Date())
            // 金曜日かどうか判定
            if(date.getDay() == 5){
              next = format(addDays(date, 3), 'yyyy-MM-dd')
            }else{
              next = format(addDays(date, 1), 'yyyy-MM-dd')
            }
          }else if(newInterval >= 1 && newInterval <= 3){ //毎週、隔週、3週毎
            next = getNextDate(next, newInterval*7)
          }else if(newInterval == 4){ //毎月
            // 曜日指定
            if(newWeek != 0){
              next = getNextMonthDate(next, newWeek, newDay);
            }else{  //日付指定
              // [todo]newDateの更新は1回でいいはずだが、初回だけ今日の日付にする必要があるので多分こうなる。
              // nextをdateオブジェクトに変更
              const date = parse(next, 'yyyy-MM-dd', new Date())
              // 日付をnewDateに上書き
              date.setDate(newDate)
              // 毎月決まった日を当番に設定する。土日祝日は考慮されない。
              next = format(addMonths(date, 1), 'yyyy-MM-dd')
            }
          }
        })
      }else{
        // memberInfos[closestLastDateIndex].orderの次のorderから順次リスケ
        // 最初はmemberInfos[closestLastDateIndex].order < memberInfo.orderをリスケ
        const lastOrder = memberInfos[closestLastDateIndex].order_number
        var nextOrder = lastOrder + 1
        
        memberInfos.forEach((memberInfo) => {
          if(memberInfo.order_number == nextOrder){
            memberInfo.next = next
            if(newInterval == 0){
              const date = parse(next, 'yyyy-MM-dd', new Date())
              // 金曜日かどうか判定
              if(date.getDay()  == 5){
                next = format(addDays(date, 3), 'yyyy-MM-dd')
              }else{
                next = format(addDays(date, 1), 'yyyy-MM-dd')
              }
            }else if(newInterval >= 1 && newInterval <= 3){ //毎週、隔週、3週毎
              next = getNextDate(next, newInterval*7)
            }else if(newInterval == 4){ //毎月
              // 曜日指定
              if(newWeek != 0){
                next = getNextMonthDate(next, newWeek, newDay);
              }else{  //日付指定
                const date = parse(next, 'yyyy-MM-dd', new Date())
                date.setDate(newDate)
                next = format(addMonths(date, 1), 'yyyy-MM-dd')
              }
            }
            nextOrder++
          }
        })
        // 次にはmemberInfos[closestLastDateIndex].order => memberInfo.orderをリスケ
        nextOrder = 1
        memberInfos.forEach((memberInfo) => {
          if(memberInfo.order_number <= lastOrder){
            memberInfo.next = next
            if(newInterval == 0){
              const date = parse(next, 'yyyy-MM-dd', new Date())
              // 金曜日かどうか判定
              if(date.getDay()  == 5){
                next = format(addDays(date, 3), 'yyyy-MM-dd')
              }else{
                next = format(addDays(date, 1), 'yyyy-MM-dd')
              }
            }else if(newInterval >= 1 && newInterval <= 3){
              next = getNextDate(next, newInterval*7)
            }else if(newInterval == 4){ //毎月
              // 曜日指定
              if(newWeek != 0){
                next = getNextMonthDate(next, newWeek, newDay);
              }else{  //日付指定
                const date = parse(next, 'yyyy-MM-dd', new Date())
                date.setDate(newDate)
                next = format(addMonths(date, 1), 'yyyy-MM-dd')
              }
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