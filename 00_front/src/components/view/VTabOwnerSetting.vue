<template>
  <div class="vtab-setting">
    <v-select
      v-model='data.nextOwner'
      :items="memberList"
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
      data: {newKey: null, nextOwner: null},
      hint: "新しい削除キーを設定してください。",
    }
  },
  props:{
    memberList:{
      type: Array,
      require: true
    }
  },
  methods:{
    GetData(){
      this.data.newKey = this.$refs.RefTextField.inputData
      return this.data
    }
  }
}
</script>
