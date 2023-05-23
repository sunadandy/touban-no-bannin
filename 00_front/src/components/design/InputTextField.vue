<template>
  <div>
    <!-- 
      @blur:フォーカスが外れたときのイベント
    -->
    <v-text-field
      variant="outlined"
      clearable
      :hint='hint'
      persistent-hint
      v-model='inputData'
    ></v-text-field>
  </div>
</template>
  
<script>
export default {
  name: "InputTextField",
  data(){
    return{
      inputData: this.initText,
    }
  },
  props:{
    hint:{
      type: String,
      require: true,
      defalt: () => ""
    },
    initText:{
      type: String,
      require: true,
      defalt: () => ""
    }
  },
}
</script>
  