<template>
  <div>
    <setting ref="RefSetting" v-bind:limited="true"/>

    <v-btn flat rounded="pill" color="primary" @click="dialog = !dialog">設定完了</v-btn>
    <v-dialog
      v-model="this.dialog"
      persistent
      width="1024"
    >
      <v-card>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <input-text-field ref="RefInputField1" v-bind:hint='hint1'/>
              </v-col>
              <v-col cols="12">
                <input-text-field ref="RefInputField2" v-bind:hint='hint2'/>
              </v-col>
              <v-col
                cols="12"
                sm="6"
              >
              </v-col>
            </v-row>
          </v-container>
          <small>*登録成功時、自動でホームに戻ります。</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="blue-darken-1"
            variant="text"
            @click="dialog = false"
          >
            閉じる
          </v-btn>
          <v-btn
            color="blue-darken-1"
            variant="text"
            @click="Register"
          >
            登録
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import Setting from '@/components/Setting'
import InputTextField from '@/components/design/InputTextField'

export default {
  name: 'CreateView',
  components: { 
    Setting,
    InputTextField,
  },
  data(){
    return {
      dialog: false,
      password: null,
      ownerNumber: null,
      hint1: "当番のオーナーの社員番号(6ケタ)を入力してください",
      hint2: "パスワードを設定してください",
      newTouban: ["id", "title", "owner", "start", "interval_type", "mail", "remind_type", "message", "password"],
      newOrder: ["toubanId", "name", "employeeNo", "order", "next", "last"],
    }
  },
  methods: {
    OwnerValidation(){
      // 入力された社員番号がアドレス帳に登録されている番号かどうかチェック
      var employeeNumberList = this.$refs.RefSetting.employeeNoList
      this.ownerNumber = this.$refs.RefInputField1.inputData
      this.password = this.$refs.RefInputField2.inputData
      for(var i=0; i<employeeNumberList.length; i++){
        if(employeeNumberList[i] == this.ownerNumber){
          // オーナーとパスワードの格納
          this.newTouban.owner = this.ownerNumber
          this.newTouban.password = this.password
          return "OK"
        }
      }
      alert("データベースに存在しない社員番号です。")
      return "NG"
    },
    Register(){
      // 設定項目チェック
      var result = {}
      result = this.OwnerValidation()
      if(result.error == "OK"){
        this.dialog = false
      }else{
        return  //途中終了する意味でのreturn
      }

      // Update時の1つ1つの処理を全て実施すればいい
      Update("option-1")
      Update("option-2")
      Update("option-3")
      Update("option-4")
      Update("option-7")

      // 当番テーブルに新規追加
      this.axios.post("/touban", this.toubanInfo)
        .then(response => {
          console.log(response)
          var toubannRefId = response.data
        }).catch(error => console.log(error))
      // 新規追加された当番テーブルIDをオーダー情報に反映
      for(var i=0; i<this.memberInfo.length;i++){
        this.memberInfo[i].toubannRefId
      }
      // オーダーテーブルに新規追加
      this.axios.post("/order", this.memberInfo)
        .then(response => {
        }).catch(error => console.log(error))

      // リダイレクト
      this.$router.push("/home")
    },
  }
}
</script>