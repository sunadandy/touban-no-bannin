<template>
  <div class="vtab-setting">
    <person-manager ref="RefPersonManager"></person-manager>
    <select
      title="メンバーを選択してください"
      class="multi"
      multiple="multiple"
      size="20"
    >
      <option v-for="item in selectable" :key="item.employeeNo">{{ item.affiliation }},{{ item.name }}</option>
      <option v-for="item in memberInfo" :key="item.employeeNo" selected>{{ item.affiliation }},{{ item.name }}</option>
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
      selectable: [], //{"affiliation", "employeeNo", "name", "email"}のJSON配列
      member: [],     //{"affiliation", "employeeNo", "name", "email"}のJSON配列
    }
  },
  props:{
    memberInfo: {
      type: Array,
      required: true,
      default: () => [],
    }
  },
  methods:{
    // selectableはJSONだが、selectedはnameのみ
    GetSelected(){
      const selected = document.getElementById("m-selected")
      var jsonArray = []
      for ( var i=0,l=selected.length; l>i; i++ ) {
        for(var j=0; j < this.selectable.length; j++){
          // 必ずどこかで一致する
          const value = selected[i].value
          const name = value.split(",")[1]
          if(name == this.selectable[j].name){
            jsonArray.push(this.selectable[j])
            break
          }
        }
      }
      return jsonArray
    },
  },
  created(){
    this.member = this.memberInfo
  },
  mounted(){
    // アドレス帳から取得したデータ（名前、社員番号、メールアドレスの連想配列）
    this.selectable = this.$refs.RefPersonManager.userData

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
  