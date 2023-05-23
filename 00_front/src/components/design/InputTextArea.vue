<template>
  <div>
    <v-textarea
      variant="outlined"
      clearable
      v-model="inputData"
    ></v-textarea>
  </div>
</template>
  
<script>
export default {
  name: "InputTextArea",
  data(){
    return{
      inputData: this.currentMessage
    }
  },
  props:{
    hint:{
      type: String,
      require: true
    },
    currentMessage: {
      type: String,
      require: true,
      default: () => "",
    },
  }
}
</script>
  