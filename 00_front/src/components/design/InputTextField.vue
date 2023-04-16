<template>
  <div>
    <v-text-field
      variant="outlined"
      clearable
      :hint='hint'
      persistent-hint
      v-model="inputData"
    ></v-text-field>
  </div>
</template>
  
<script>
export default {
  name: 'InputTextField',
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
  