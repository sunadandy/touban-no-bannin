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
      inputData: "",
    }
  },
  props:{
    hint:{
      type: String,
      require: true
    }
  }
}
</script>
  