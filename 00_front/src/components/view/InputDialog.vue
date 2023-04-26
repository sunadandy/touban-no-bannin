<template>
  <div>
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

export default {
  name: 'InputDialog',
  components: { 
    InputTextField,
  },
  data(){
    return {
      dialog: false,
      password: "",
      ownerNumber: "",
      hint1: "当番のオーナーの社員番号を入力してください",
      hint2: "パスワードを設定してください",
    }
  },
  props:{
    selected:{
      type: Array,
      required: true,
      default: () => [],
    }
  },
  methods: {
    OpenDialog(){
      this.dialog = true
    },
    Validation(){
      const ownerNumber = this.$refs.RefInputField1.inputData
      const password = this.$refs.RefInputField2.inputData
      var owner = new String()
      
      // 入力された社員番号がアドレス帳に登録されている番号かどうかチェック
      for(var i=0; i<this.selected.length; i++){
        if(this.selected[i].emplyoeeNo == ownerNumber){
          owner = this.selected[i].name
          this.ownerNumber = ownerNumber
          this.password = password
          this.dialog = false
          
          this.$emit('CloseDialog', {password: this.password, owner: owner})
          return true //照合が成功したら処理を抜けたいのでtrueに意味はないがreturn
        }
      }
      alert("オーナーが選択メンバーに含まれていません。")
      this.dialog = true
    }
  }
}
</script>