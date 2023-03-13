import  SelectBasicExample  from '../../components/Shares/search/search__form.js'
import  SelectInput from '../../components/Shares/search/search__input.js'
import  BasicExample from '../../components/Shares/table/table.js'
import './topic.css'


function Topic() {
  var openSearch = false
    return (
      <>
      
        <div id ="topic__name">
          <p>Đề tài nghiên cứu khoa học</p>
          <SelectInput></SelectInput>
        </div>
        <div id ="topic__search">
          <SelectBasicExample openSearch={openSearch}></SelectBasicExample>
        </div>
        <div id = "topic__table">
          <BasicExample></BasicExample>
        </div>
      
      </>
    );
  }
  
  export default Topic;