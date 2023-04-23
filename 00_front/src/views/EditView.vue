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
      // 更新前チェックとデータの取得
      var result = this.$refs.RefSetting.Validation()
      if(result.status == "OK"){
        result.data
      // }else{
        
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