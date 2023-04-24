<template>
  <div>
    <setting ref="RefSetting" v-bind:limited="true"/>
    
    <v-btn flat rounded="pill" color="primary" @click="Register">設定完了</v-btn>
    <input-dialog ref="RefDialog"/>
  </div>
</template>

<script>
import Setting from '@/components/Setting'
import InputDialog from '@/components/view/InputDialog'

export default {
  name: 'CreateView',
  components: { 
    Setting,
    InputDialog,
  },
  data(){
    return {
      dialog: false,
      newTouban: ["id", "title", "owner", "start", "interval_type", "mail", "remind_type", "message", "password"],
      newOrder: ["toubanId", "name", "employeeNo", "order", "next", "last"],
    }
  },
  methods: {
    Register(){
      const result = this.$refs.RefDialog.OpenDialog()
      console.log(result)
      // // 当番テーブルに新規追加
      // this.axios.post("/touban", this.toubanInfo)
      //   .then(response => {
      //     console.log(response)
      //     var toubannRefId = response.data
      //   }).catch(error => console.log(error))
      // // 新規追加された当番テーブルIDをオーダー情報に反映
      // for(var i=0; i<this.memberInfo.length;i++){
      //   this.memberInfo[i].toubannRefId
      // }
      // // オーダーテーブルに新規追加
      // this.axios.post("/order", this.memberInfo)
      //   .then(response => {
      //   }).catch(error => console.log(error))
    },
  }
}
</script>