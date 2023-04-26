<template>
  <div>
    <calc-member-info ref="RefCaclcMemberInfo"/>
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
import CalcMemberInfo from '@/components/module/CalcMemberInfo'

export default {
  name: 'EditView',
  components: { 
    Setting,
    CalcMemberInfo,
  },
  data(){
    return {
      // current_toubanInfo:{},
      // current_memberInfo:[],
    }
  },
  methods:{
    Update(){
      // 子(Setting)に入力項目のチェックを要求
      var result = this.$refs.RefSetting.Validation()
      if(!result.status){
        alert("${currentTab}の設定が正しくありません")
        return
      }
      
      // データベース更新(アクティブタブ単位でデータを取得)
      switch(result.tab){
        // 当番名
        case "option-1":
          var current_toubanInfo = {}
          current_toubanInfo.title = result.data
          this.axios.put("/touban", this.current_toubanInfo)
          .then(response => {
            console.log(response)
          }).catch(error => console.log(error))
          break
        // メンバー編集
        case "option-2":
          const toubanId = this.$route.params.id
          var selected = result.data
          var new_memberInfos = this.$refs.RefCaclcMemberInfo.GetNewMemberInfo(toubanId, selected)
          // 古いメンバー情報を削除
          this.axios.delete("/member", {params:{id: toubanId}})
          .then(response => {
            // 新しいメンバー情報を登録
            this.axios.post("/member", {new_memberInfos})
          }).catch(error => console.log(error))
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
    },
  },
  mounted(){
    const toubanId = this.$route.params.id
    const toubanInfo = this.$store.getters.GetToubanByID(toubanId)
    this.current_toubanInfo = toubanInfo[0]  //GetToubanByIDがレコードフィルタ結果を配列で返してくる
  }
}
</script>