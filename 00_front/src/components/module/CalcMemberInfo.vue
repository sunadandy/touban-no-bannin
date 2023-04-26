<template>
  <div>
  </div>
</template>

<script>
import { addDays, format, parseISO, startOfToday } from 'date-fns'

export default {
  name: 'CalcMemberInfo',
  data(){
    return{
      toubanInfo: [],
      interval_type: 2  //デフォルト値
    }
  },
  methods:{
    GetNewMemberInfo(toubanId, selected){
      var new_memberInfos = []

      // 当番IDに一致する順番レコードを取得
      var current_memberInfos = []
      current_memberInfos = this.$store.getters.GetMemberByToubanId(toubanId)
      // CreateかUpdateかを取得したレコードの長さで判定
      if(current_memberInfos.length == 0){
        var currentMaxOrder = 0
        var currentMaxNextDate = format(startOfToday(), 'yyyy-MM-dd')
        for(var i=0; i<selected.length; i++){
          var memberInfo = {}
          memberInfo.touban_id = toubanId
          memberInfo.name = selected[i].name
          memberInfo.emplyoee_number = "E" + selected[i].emplyoeeNo   //バックエンド側で数値文字列を上手く扱えないのでE付きに。
          memberInfo.order_number = currentMaxOrder + 1
          memberInfo.last = "2100-1-1"  //適当な初期値
          memberInfo.next = this.CalcNextDate(currentMaxNextDate)
          memberInfo.email = selected[i].email
          new_memberInfos.push(memberInfo)

          currentMaxOrder = memberInfo.order_number
          currentMaxNextDate = memberInfo.next
        }
      }else{
        // 現在のorderの最大値を取得
        var currentMaxOrder = current_memberInfos.reduce((prev, current) => (prev.order > current.order) ? prev.order : current.order);
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
              memberInfo.touban_id = this.toubanId
              memberInfo.name = selected[j].name
              memberInfo.emplyoee_number = "E" + selected[j].emplyoeeNo   //バックエンド側で数値文字列を上手く扱えないのでE付きに。
              memberInfo.order_number = currentMaxOrder + 1
              memberInfo.last = "2100-1-1"  //適当な初期値
              memberInfo.next = this.CalcNextDate(currentMaxNextDate)
              memberInfo.mail = selected[j].email
              new_memberInfos.push(memberInfo)
  
              currentMaxOrder = memberInfo.order
              currentMaxNextDate = memberInfo.next
            }
          }
        }
      }
      return new_memberInfos
    },
    CalcNextDate(currentMaxNextDate){
      const date = parseISO(currentMaxNextDate)
      switch(this.interval_type){
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
  created(){
    const toubanId = this.$route.params.id
    // Editモードの場合はURLからID取得
    if(toubanId != undefined){
      // toubanIDに一致するレコード取得
      this.toubanInfo = this.$store.getters.GetToubanByID(toubanId)
      this.interval_type = int(this.toubanInfo[0].interval_type)
    }
  }
}
</script>