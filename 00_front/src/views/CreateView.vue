<template>
  <div>
    <setting ref="RefSetting" v-bind:limited="true"/>
    <calc-member-info ref="RefCaclcMemberInfo"/>
    
    <v-btn flat rounded="pill" color="primary" @click="Validation">設定完了</v-btn>
    <v-dialog v-model="dialog" persistent width="1024">
      <input-dialog v-bind:hints="hints" @clickSubmit="onSubmit" @clickClose="onClose"/>
    </v-dialog>
  </div>
</template>

<script>
import Setting from '@/components/Setting'
import InputDialog from '@/components/design/InputDialog'
import CalcMemberInfo from '@/components/module/CalcMemberInfo'

// グローバル変数
var result  //本当は嫌だが、こうしないとRegisterに渡せない。dataプロパティでもいいがresultはあくまで一時変数に過ぎない

export default {
  name: 'CreateView',
  data(){
    return{
      dialog: false,
      hints: ["当番のオーナーの社員番号を入力してください", "パスワードを設定してください"]
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
      if(!result.status){
        alert("いずれかの入力に不備があります。")
        return false
      }
      // オーナーとパスワードの設定
      this.dialog = true
    },
    Register(owner, password){
      var newTouban = {}
      var newMember = {}
      const maxToubanId = this.$store.state.currenToubanTable.reduce((prev, current) => {
        return (prev.id > current.id) ? prev.id : current.id
      }, 0);
      const nextToubanId = maxToubanId + 1
      // オーナーのメールアドレス取得
      const ownerData = result.data[1].data.find(data => data.name == owner);
      const email = ownerData ? ownerData.email : "";

      newTouban.id = nextToubanId
      newTouban.title = result.data[0].data
      newTouban.owner = owner
      newTouban.start = result.data[2].data.startDate
      newTouban.interval_type = parseInt(result.data[2].data.interval)
      newTouban.mailing = result.data[4].data.mailing
      newTouban.timing = result.data[4].data.timing
      newTouban.message = result.data[3].data
      newTouban.password = password
      newTouban.cc = email + ";" + result.data[4].data.cc   //オーナーはCCに自動追加
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
    onClose(){
      this.dialog = false
    },
    onSubmit(payload) {
      const ownerNumber = payload[0]
      const password = payload[1]
      const selected = this.$refs.RefSetting.GetSelected()

      // パスワードの空白チェック
      if(password == ""){
        alert("パスワードが空白です。")
        return false
      }

      // 入力された社員番号がアドレス帳に登録されている番号かどうかチェック
      for(var i=0; i<selected.length; i++){
        if(selected[i].emplyoeeNo == ownerNumber){
          const owner = selected[i].name
          this.Register(owner, password)
          this.dialog = false
          return true   //これがないと必ずalertを実行してしまうので必須
        }
      }
      alert("オーナーが選択メンバーに含まれていません。")
    },
  }
}
</script>