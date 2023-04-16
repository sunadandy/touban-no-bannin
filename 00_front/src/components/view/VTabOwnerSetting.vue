<template>
  <div class="vtab-setting">
    <v-select
      v-model="nextOwner"
      :items="memberList"
      hint="オーナー変更先を選んでください"
      persistent-hint
      variant="outlined"
    ></v-select>
    <input-text-field v-bind:hint='hint'/>
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
      nextOwner: "",
      newkey: "",
      hint: "新しい削除キーを設定してください。",
    }
  },
  props:{
    memberList:{
      type: Array,
      require: true
    }
  }
}
</script>
