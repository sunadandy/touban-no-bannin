<template>
  <div class="vtab-setting">
    <person-manager ref="RefPersonManager"></person-manager>
    <select
      title="メンバーを選択してください"
      class="multi"
      multiple="multiple"
      size="20"
    >
      <option v-for="item in selectable" :key="item.employeeNo">{{ item.name }}</option>
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
      selectable: [],
    }
  },
  methods:{
    // selectableと選択されたアイテム(name)を比較して一致する要素を返す
    GetSelected(){
      const selected = document.getElementById("m-selected")
      var array = []
      for ( var i=0,l=selected.length; l>i; i++ ) {
        for(var j=0; j < this.selectable.length; j++){
          if(selected[i].value == this.selectable[j].name){
            array.push(this.selectable[j])
            break
          }
        }
      }
      return array
    },
  },
  mounted(){
    // アドレス帳から取得したデータ（名前、社員番号、メールアドレスの連想配列）
    this.selectable = this.$refs.RefPersonManager.userData
    
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
  