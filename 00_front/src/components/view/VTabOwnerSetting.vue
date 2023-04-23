<template>
  <div class="vtab-setting">
    <v-select
      v-model='data.nextOwner'
      :items="memberList.name"
      hint="オーナー変更先を選んでください"
      persistent-hint
      variant="outlined"
    ></v-select>
    <!-- <input-text-field v-bind:hint='hint' @CbkChangeText="CbkChangeText"/> -->
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
      NameOnlyList: [],
      data: {newPassword: null, nextOwner: null},
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
      this.data.newPassword = this.$refs.RefTextField.inputData
      return this.data
    }
  }
}
</script>
