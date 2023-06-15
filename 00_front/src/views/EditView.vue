<template>
  <div>
    <calc-member-info ref="RefCaclcMemberInfo"/>
    <setting ref="RefSetting"/>
    <v-btn flat rounded="pill" color="primary" @click="onClick">更新</v-btn>
    <v-dialog v-model="dialog" persistent width="600">
      <input-dialog v-bind:hints="hints" @clickSubmit="onSubmit" @clickClose="onClose"/>
    </v-dialog>
  </div>
</template>

<script>
import Setting from '@/components/Setting'
import CalcMemberInfo from '@/components/module/CalcMemberInfo'
import InputDialog from '@/components/design/InputDialog'

export default {
  name: 'EditView',
  components: { 
    Setting,
    CalcMemberInfo,
    InputDialog,
  },
  data(){
    return {
      toubanId: -1,
      dialog: false,
      hints: ["パスワードを入力してください"],
      current_toubanInfo: {},
      current_memberInfos: [],
      result: {},
    }
  },
  methods:{
    onClick(){
      // 子(Setting)に入力項目のチェックを要求
      this.result = this.$refs.RefSetting.Validation()
      // status=falseなら終了
      if(!this.result.status){
        // alert(`${result.tab}の設定が正しくありません`)
        return
      }

      // selectedにオーナーが含まれていない場合はNG
      if(this.result.tab == "option-2"){
        const foundOwner = this.result.data.some((item) => {
          return item.name.includes(this.current_toubanInfo.owner)
        })
        if(!foundOwner){
          alert("オーナーが選択されていません")
          return false
        }
      }
      this.dialog = true
    },
    onClose(){
      this.dialog = false
    },
    onSubmit(){
      // データベース更新(アクティブタブ単位でデータを取得)
      switch(this.result.tab){
        // 当番名
        case "option-1":
          var updateToubanInfo = this.current_toubanInfo
          updateToubanInfo.title = this.result.data
          this.axios.put("/touban", updateToubanInfo)
          .then(response => {
            console.log(response)
          }).catch(error => console.log(error))
          break
        // メンバー編集
        case "option-2":
          const selected = this.result.data
          // メンバーの更新
          var updateMemberInfos = this.$refs.RefCaclcMemberInfo.ReplaceMember(this.current_memberInfos, selected)
          // 変更したメンバーの次回実施予定日設定
          const scheduling = this.current_toubanInfo.scheduling
          var data = scheduling.split("-")
          var interval = parseInt(data[0])
          const day = parseInt(data[1])
          var nextDate = this.current_toubanInfo.next
          updateMemberInfos = this.$refs.RefCaclcMemberInfo.ReSchedule(updateMemberInfos, interval, day, nextDate)
          // 古いメンバー情報を削除
          this.axios.delete("/member", {params:{id: this.toubanId}})
          .then(response => {
            // 新しいメンバー情報を登録
            this.axios.post("/member", updateMemberInfos)
          }).catch(error => console.log(error))
          break
        // スケジュール設定
        case "option-3":
          // 当番情報の更新
          var updateToubanInfo = this.current_toubanInfo
          const newScheduling = this.result.data.scheduling
          const newStartDate = this.result.data.startDate
          updateToubanInfo.scheduling = newScheduling
          updateToubanInfo.next = newStartDate
          // 次回実施日のリスケ
          var data = newScheduling.split("-")
          const newInterval = parseInt(data[0])
          const newDay = parseInt(data[1])
          var updateMemberInfos = this.$refs.RefCaclcMemberInfo.ReSchedule(this.current_memberInfos.sort((a, b) => a.order_number - b.order_number), newInterval, newDay, newStartDate)
          this.axios.put("/touban", updateToubanInfo)
          .then(response => {
            this.axios.put("/member", updateMemberInfos)
            .then(response => {
              console.log(response)
            }).catch(error => console.log(error))
          }).catch(error => console.log(error))
          break
        // メッセージ設定
        case "option-4":
          var updateToubanInfo = this.current_toubanInfo
          updateToubanInfo.message = this.result.data
          this.axios.put("/touban", updateToubanInfo)
          .then(response => {
            console.log(response)
          }).catch(error => console.log(error))
          break
          // 順番変更
        case "option-5":
          var newOrder = this.result.data
          newOrder.forEach((element, index) => {
            element.order_number = index + 1
          });
          // 次回実施日のリスケ
          var sch = this.current_toubanInfo.scheduling.split("-")
          var interval = parseInt(sch[0])
          var date = parseInt(sch[1])
          var nextDate = this.current_toubanInfo.next
          var updateMemberInfos = this.$refs.RefCaclcMemberInfo.ReSchedule(newOrder, interval, date, nextDate)
          this.axios.put("/member", updateMemberInfos)
          .then(response => {
            console.log(response)
          }).catch(error => console.log(error))
          break
        // オーナー変更
        case "option-6":
          var updateToubanInfo = this.current_toubanInfo
          updateToubanInfo.owner = this.result.data.owner
          updateToubanInfo.password = this.result.data.password
          updateToubanInfo.cc = updateToubanInfo.cc + ";" + this.result.data.email
          this.axios.put("/touban", updateToubanInfo)
          .then(response => {
            console.log(response)
          }).catch(error => console.log(error))
          break
        // メール配信設定
        case "option-7":
          var updateToubanInfo = this.current_toubanInfo
          updateToubanInfo.mailing = this.result.data.mailing
          updateToubanInfo.timing = this.result.data.timing
          updateToubanInfo.cc = this.result.data.cc
          this.axios.put("/touban", updateToubanInfo)
          .then(response => {
            console.log(response)
          }).catch(error => console.log(error))
          break
      }
      this.dialog = false
    },
  },
  created(){
    this.toubanId = this.$route.params.id
    const toubanInfo = this.$store.getters.GetToubanByID(this.toubanId)
    this.current_toubanInfo = toubanInfo[0]  //GetToubanByIDがレコードフィルタ結果を配列で返してくる
    this.current_memberInfos = this.$store.getters.GetMemberByToubanId(this.toubanId)
  }
}
</script>