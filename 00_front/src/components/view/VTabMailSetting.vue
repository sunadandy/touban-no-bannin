<template>
  <div class="vtab-setting">
    <div>
      ここではメールの配信設定が行えます。<br>
      デフォルト設定では、次回担当者に実施日前日にメールが配信されるようになっています。<br>
    </div>
    <div>
      <v-switch
        v-model="model"
        hide-details
        :label="`メール配信: ${model.toString()}`"
      ></v-switch>
      <v-select
      v-if="model == true"
        hint="何日前に配信しますか？"
        persistent-hint
        :items=[1,2,3,4,5,6,7]
        variant="outlined"
        v-model='remindDate'
      ></v-select>
      <input-text-field v-bind:hint='hint' ref="RefInputField" v-if="model == true"/>
    </div>
  </div>
</template>

<script>
import InputTextField from '@/components/design/InputTextField'

export default {
  name: 'MailSetting',
  components:{
    InputTextField
  },
  data(){
    return{
      model: true,
      hint: "CC設定(複数の場合はコロン区切り)",
      remindDate: 1
    }
  },
  methods:{
    GetMailSetting(){
      if(this.model === false){
        return [this.model, null, null]
      }else{
        var cc = this.$refs.RefInputField.inputData
        console.log(cc)
        return [this.model, this.remindDate, cc]
      }
    }
  }
}
</script>
