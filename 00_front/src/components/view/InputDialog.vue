<template>
  <div>
    <person-manager ref="RefPersonManager"></person-manager>

    <v-dialog
      v-model=this.dialog
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
            @click="Validation"
          >
            登録
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import InputTextField from '@/components/design/InputTextField'
import PersonManager from '@/components/model/PersonManager'

export default {
  name: 'InputDialog',
  components: { 
    InputTextField,
    PersonManager,
  },
  data(){
    return {
      dialog: false,
      password: null,
      ownerNumber: null,
      hint1: "当番のオーナーの社員番号(6ケタ)を入力してください",
      hint2: "パスワードを設定してください",
    }
  },
  methods: {
    OpenDialog(){
      this.dialog = true
    },
    Validation(){
      // 入力された社員番号がアドレス帳に登録されている番号かどうかチェック
      var userData = this.$refs.RefPersonManager.userData
      var ownerNumber = this.$refs.RefInputField1.inputData
      var password = this.$refs.RefInputField2.inputData
      for(var i=0; i<userData.length; i++){
        if(userData[i].emplyoeeNo == ownerNumber){
          // オーナーとパスワードの格納
          this.ownerNumber = ownerNumber
          this.password = password
          this.dialog = false

          return true
        }
      }
      alert("データベースに存在しない社員番号です。")
      this.dialog = true
      return false
    }
  }
}
</script>