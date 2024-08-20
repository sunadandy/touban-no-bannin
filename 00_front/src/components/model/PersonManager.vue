<template>
  <div>
  </div>
</template>
  
<script>
import csvData from "@/assets/部内リスト.csv"

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
      item["affiliation"] = this.csvData[i][4]
      item["emplyoeeNo"] = this.csvData[i][2]
      item["name"] = this.csvData[i][0]
      item["email"] = this.csvData[i][3]
      this.userData.push(item)
    }
  }
};
</script>