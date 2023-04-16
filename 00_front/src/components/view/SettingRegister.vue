<template>
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
              <input-text-field v-bind:hint='hint1'/>
            </v-col>
            <v-col cols="12">
              <input-text-field v-bind:hint='hint2'/>
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
</template>

<script>
import InputTextField from '@/components/design/InputTextField.vue'

export default {
  name: 'SettingRegister',
  components:{
    InputTextField
  },
  data(){
    return {
      dialog: false,
      deleteKey: [],
      employeeNo: [],
      hint1: "当番のオーナーの社員番号(6ケタ)を入力してください",
      hint2: "削除キーを設定してください",
    }
  },
  methods: {
    Register(){
      this.$emit('CbkRegister', {deleteKey: this.deleteKey, owner: this.employeeNo})
    }
  }
}
</script>