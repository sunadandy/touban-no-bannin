<template>
  <div>
    <calc-member-info ref="RefCaclcMemberInfo"/>
    <setting ref="RefSetting"/>
    <v-btn flat rounded="pill" color="primary" @click="Update">更新</v-btn>
    <v-overlay :model-value="overlay" class="align-center justify-center">
      <v-progress-circular
        color="primary"
        indeterminate
        size="64"
      ></v-progress-circular>
    </v-overlay>
  </div>
</template>

<script>
import Setting from '@/components/Setting'
import CalcMemberInfo from '@/components/module/CalcMemberInfo'

export default {
  name: 'EditView',
  components: { 
    Setting,
    CalcMemberInfo,
  },
  data(){
    return {
      toubanId: -1,
      current_toubanInfo: {},
      currentOwner: "",
      overlay: false,
    }
  },
  watch: {
    overlay (val) {
      val && setTimeout(() => {
        this.overlay = false
      }, 500)
    },
  },
  methods:{
    Update(){
      // 子(Setting)に入力項目のチェックを要求
      var result = this.$refs.RefSetting.Validation()
      // status=falseなら終了
      if(!result.status){
        // alert(`${result.tab}の設定が正しくありません`)
        return
      }

      // selectedにオーナーが含まれていない場合はNG
      if(result.tab == "option-2"){
        const foundOwner = result.data.some((item) => {
          return item.name.includes(this.currentOwner)
        })
        if(!foundOwner){
          alert("オーナーが選択されていません")
          return false
        }
      }
      
      // データベース更新(アクティブタブ単位でデータを取得)
      switch(result.tab){
        // 当番名
        case "option-1":
          var updateToubanInfo = this.current_toubanInfo
          updateToubanInfo.title = result.data
          this.axios.put("/touban", updateToubanInfo)
          .then(response => {
            console.log(response)
          }).catch(error => console.log(error))
          break
        // メンバー編集
        case "option-2":
          const selected = result.data
          const interval_type = this.current_toubanInfo.interval_type
          const updateMemberInfos = this.$refs.RefCaclcMemberInfo.CalcuNewMemberInfo(this.toubanId, interval_type, selected)
          // 古いメンバー情報を削除
          this.axios.delete("/member", {params:{id: this.toubanId}})
          .then(response => {
            // 新しいメンバー情報を登録
            this.axios.post("/member", updateMemberInfos)
          }).catch(error => console.log(error))
          break
        // スケジュール設定
        case "option-3":
          break
        // メッセージ設定
        case "option-4":
          var updateToubanInfo = this.current_toubanInfo
          updateToubanInfo.message = result.data
          this.axios.put("/touban", updateToubanInfo)
          .then(response => {
            console.log(response)
          }).catch(error => console.log(error))
          break
          // 順番変更
        case "option-5":
          var newOrder = result.data
          newOrder.forEach((element, index) => {
            element.order_number = index + 1
          });
          this.axios.put("/member", newOrder)
          .then(response => {
            console.log(response)
          }).catch(error => console.log(error))
          break
        // オーナー変更
        case "option-6":
          var updateToubanInfo = this.current_toubanInfo
          updateToubanInfo.owner = result.data.owner
          updateToubanInfo.password = result.data.password
          updateToubanInfo.cc = updateToubanInfo.cc + ";" + result.data.email
          this.axios.put("/touban", updateToubanInfo)
          .then(response => {
            console.log(response)
          }).catch(error => console.log(error))
          break
        // メール配信設定
        case "option-7":
          var updateToubanInfo = this.current_toubanInfo
          updateToubanInfo.mailing = result.data.mailing
          updateToubanInfo.timing = result.data.timing
          updateToubanInfo.cc = result.data.cc
          this.axios.put("/touban", updateToubanInfo)
          .then(response => {
            console.log(response)
          }).catch(error => console.log(error))
          break
      }
      // 更新を視覚的に通知するための演出
      this.overlay = true
    },
  },
  created(){
    this.toubanId = this.$route.params.id
    const toubanInfo = this.$store.getters.GetToubanByID(this.toubanId)
    this.current_toubanInfo = toubanInfo[0]  //GetToubanByIDがレコードフィルタ結果を配列で返してくる
    this.currentOwner = this.current_toubanInfo.owner
  }
}
</script>