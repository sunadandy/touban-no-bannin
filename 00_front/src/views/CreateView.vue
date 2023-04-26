<template>
  <div>
    <setting ref="RefSetting" v-bind:limited="true"/>
    <calc-member-info ref="RefCaclcMemberInfo"/>
    
    <v-btn flat rounded="pill" color="primary" @click="Validation">設定完了</v-btn>
    <input-dialog ref="RefDialog" v-bind:selected='selected' @CloseDialog="Register"/>
  </div>
</template>

<script>
import Setting from '@/components/Setting'
import InputDialog from '@/components/view/InputDialog'
import CalcMemberInfo from '@/components/module/CalcMemberInfo'

// グローバル変数
var result  //本当は嫌だが、こうしないとRegisterに渡せない。dataプロパティでもいいがresultはあくまで一時変数に過ぎない

export default {
  name: 'CreateView',
  data(){
    return{
      selected:[],
    }
  },
  components: { 
    Setting,
    InputDialog,
    CalcMemberInfo,
  },
  methods: {
    Validation(){
      // 子コンポーネント(setting)にバリデーション要求
      result = this.$refs.RefSetting.Validation()
      if(result.status){
        // オーナーとパスワードの設定
        this.selected = this.$refs.RefSetting.GetSelected()
        this.$refs.RefDialog.OpenDialog()
      }else{
        alert("いずれかの入力に不備があります。")
      }
    },
    Register(payload){
      var newTouban = {}
      var newMember = {}
      const maxToubanId = this.$store.state.currenToubanTable.reduce((prev, current) => {
        return (prev.id > current.id) ? prev.id : current.id
      }, 0);
      const nextToubanId = maxToubanId + 1
      const owner = payload.owner
      const password = payload.password

      newTouban.id = nextToubanId
      newTouban.title = result.data[0].data
      newTouban.owner = owner
      newTouban.start = result.data[2].data.startDate
      newTouban.interval_type = parseInt(result.data[2].data.interval)
      newTouban.mailing = result.data[4].data.mailing
      newTouban.timing = result.data[4].data.timing
      newTouban.message = result.data[3].data
      newTouban.password = password
      newTouban.cc = result.data[4].data.cc
      newMember = this.$refs.RefCaclcMemberInfo.CalcuNewMemberInfo(nextToubanId, newTouban.interval_type, result.data[1].data)

      // 当番テーブルに新規追加
      this.axios.post("/touban", newTouban)
      .then(response => {
        // メンバーテーブルに新規追加
        this.axios.post("/member", newMember)
          .then(response => {
          }).catch(error =>
            // メンバーの追加に失敗したら当番を削除
            this.axios.delete("/touban", {params:{id: nextToubanId}})
          )
      }).catch(error => console.log(error))      


      this.$router.push("/home")
    },
  }
}
</script>