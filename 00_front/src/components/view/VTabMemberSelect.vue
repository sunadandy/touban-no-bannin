<template>
  <div class="vtab-setting">
    <person-manager ref="RefPersonManager"></person-manager>
    <select
      title="メンバーを選択してください"
      class="multi"
      multiple="multiple"
      size="20"
    >
      <option v-for="item in selectable" :key="item.employeeNo">{{ item.name }}, {{ item.affiliation }}</option>
      <option v-for="item in currentMemberInfo" :key="item.employeeNo" selected>{{ item.name }}, {{ item.affiliation }}</option>
    </select>
  </div>
</template>

<script>
// import jQuery from 'jquery' //インポートししなくていいっぽい
import PersonManager from '@/components/model/PersonManager'

export default {
  name: 'MemberSelect',
  components:{
    PersonManager,
  },
  data(){
    return{
      userData: [], //CSVからのオリジナルデータ。変更してはいけない
      selectable: [], //{"affiliation", "employeeNo", "name", "email"}のJSON配列
      member: [], //{"affiliation", "employeeNo", "name", "email"}のJSON配列
    }
  },
  props:{
    currentMemberInfo: {
      type: Array,
      required: true,
      default: () => [],
    }
  },
  methods:{
    GetSelected(){
      const selected = document.getElementById("m-selected")
      var jsonArray = []
      const userNameElementNumber = 0
      for ( var i=0,l=selected.length; l>i; i++ ) {
        for(var j=0; j < this.userData.length; j++){
          // 必ずどこかで一致する
          const value = selected[i].value
          const name = value.split(",")[userNameElementNumber]
          if(name == this.userData[j].name){
            jsonArray.push(this.userData[j])
            break
          }
        }
      }
      return jsonArray
    },
  },
  created(){
    this.member = this.currentMemberInfo.map(member => {
      return {
        name: member.name,
        affiliation: member.affiliation,
        email: member.emial,
        emplopeeNo: member.employeeNo,
      }
    })
  },
  mounted(){
    // アドレス帳から取得したデータ（名前、社員番号、メールアドレスの連想配列）
    this.userData = this.$refs.RefPersonManager.userData
    this.selectable = this.userData.concat()    //複製

    // 親コンポーネントから取得したcurrentMemberInfoと照合して一致する名前をselectableから削除
    if(this.member.length != 0){
      for(var i=0; i < this.selectable.length; i++){
        for(var j=0; j < this.member.length; j++){
          if(this.selectable[i].name == this.member[j].name){
            this.selectable.splice(i, 1)
            i-- // 削除された要素のインデックスを調整
            break
          }
        }
      }
    }
  
    // DOMの更新を待ってから処理。nextTickがないとメンバーが表示されない
    this.$nextTick(function() {
      // MultiSelectedTableのマウント
      $('.multi').multiselectable({
        selectableLabel: '選択可能一覧',
        selectedLabel: '選択済み',
      });
    });
  },
}
</script>
  