<template>
  <div class="vtab-setting">
    <v-select
      v-model="owner"
      :items="nameList"
      hint="オーナー変更先を選んでください"
      persistent-hint
      variant="outlined"
    ></v-select>
    <input-text-field v-bind:hint='hint' ref="RefTextField"/>
  </div>
</template>

<script>
import InputTextField from '@/components/design/InputTextField.vue'

export default {
  name: 'OwnerSetting',
  components:{
    InputTextField,
  },
  data(){
    return{
      owner: null,
      nameList: this.memberList.map(member => member.name),
      hint: "新しいパスワードを設定してください。",
    }
  },
  props:{
    memberList:{
      type: Array,
      required: true,
      default: () => [],
    }
  },
  methods:{
    GetData(){
      const password = this.$refs.RefTextField.inputData
      var email = ""
      if(this.owner != null){
        email = this.memberList.find(member => member.name == this.owner).email
      }
      return {
        owner: this.owner,
        password: password,
        email: email,
      }
    }
  }
}
</script>
