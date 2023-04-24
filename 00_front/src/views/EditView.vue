<template>
  <div>
    <setting
      v-bind:current_toubanInfo="current_toubanInfo"
      v-bind:current_memberInfo="current_memberInfo"
      ref="RefSetting"
    />
    <v-btn flat rounded="pill" color="primary" @click="Update">更新</v-btn>
  </div>
</template>

<script>
import Setting from '@/components/Setting'

export default {
  name: 'CreateView',
  components: { 
    Setting,
  },
  data(){
    return {
      current_toubanInfo:[],
      current_memberInfo:[],
    }
  },
  methods:{
    Update(){
      var currentTab = this.$refs.RefSetting.tab
      // 更新前チェックとデータの取得
      var result = this.$refs.RefSetting.Validation()
      if(result.status == "OK"){
        result.data
      }else{
        alert("${currentTab}の設定が正しくありません")
      }

      // データベース更新
      switch(tab){
        // 当番名
        case "option-1":
          this.axios.post("/touban", data)
          .then(response => {
            console.log(response)
          }).catch(error => console.log(error))
          break
        // メンバー編集
        case "option-2":
          break
        // スケジュール設定
        case "option-3":
          break
        // メッセージ設定
        case "option-4":
          break
          // 順番変更
        case "option-5":
          break
        // オーナー変更
        case "option-6":
          break
        // メール配信設定
        case "option-7":
          break
      }
    }
  },
  // mounted(){
  //   var urlPathId = this.$route.params.id
  //   this.axios.get("/order", {params:{id: urlPathId}})
  //     .then(response => {
  //       this.member = response.data
  //     })
  //     // .catch(error => console.log(error))
  // }
}
</script>