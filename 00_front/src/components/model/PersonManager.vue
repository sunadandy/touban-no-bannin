<template>
  <div>
  </div>
</template>
  
<script>
import csvData from "@/assets/address.csv"

export default {
  name: 'PersonManager',
  data() {
    return {
      csvData: csvData,
      userData: [],
    };
  },
  mounted(){
    for(var i=0, l=this.csvData.length; i < l; i++){
      var item = {}
      item["emplyoeeNo"] = this.csvData[i][3]
      item["name"] = this.csvData[i][4]
      item["email"] = this.csvData[i][6]
      this.userData.push(item)
    }
  }
};
</script>