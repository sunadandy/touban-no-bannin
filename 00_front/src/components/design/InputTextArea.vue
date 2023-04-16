<template>
  <div class="message-form">
    <v-textarea
      name="input-7-1"
      :hint='hint'
      persistent-hint
      variant="outlined"
      auto-grow
    ></v-textarea>
  </div>
</template>
  
<script>
export default {
  name: 'InputTextArea',
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
  