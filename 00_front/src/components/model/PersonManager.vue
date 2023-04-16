<template>
  <div>
  </div>
</template>
  
<script>
import memberInfo from "@/assets/address.csv"

export default {
  name: 'PersonManager',
  data() {
    return {
      emplyoeeNo: [],   //社員番号
      name: [],         //氏名
      address: [],      //メールアドレス
      memberInfo: memberInfo,
    };
  },
  methods: {
    Members(){
      return this.name
    },
    EmployeeNo(){
      return this.emplyoeeNo
    }
  },
  mounted(){
    for(var i=0, l=this.memberInfo.length; i < l; i++){
      this.emplyoeeNo.push(memberInfo[i][3])
      this.name.push(memberInfo[i][4])
      this.address.push(memberInfo[i][5])
    }
  }
};
</script>