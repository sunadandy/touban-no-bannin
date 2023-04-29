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
        v-model='timing'
      ></v-select>
      <input-text-field v-bind:hint='hint' v-bind:initText="currentCc" ref="RefInputField" v-if="model == true"/>
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
    return {
      model: this.isValid,
      hint: "CC設定(複数の場合はコロン区切り)。オーナーは自動でCCに追加されます",
      timing: 1,
    }
  },
  props:{
    currentCc:{
      type: String,
      required: true,
      default: () => "",
    },
    isValid:{
      type: Boolean,
      required: true,
      default: () => true,
    }
  },
  methods:{
    GetMailSetting(){
      if(this.model == false){
        return {
          mailing: this.model,     //MYSQLでのbooleanがtinyintなので0/1に変換
          timing: null,  //デフォルト値
          cc: null,
        }
      }else{
        var cc = this.$refs.RefInputField.inputData
        return {
          mailing: this.model,   //MYSQLでのbooleanがtinyintなので0/1に変換
          timing: this.timing,
          cc: cc,
        }
      }
    }
  }
}
</script>
