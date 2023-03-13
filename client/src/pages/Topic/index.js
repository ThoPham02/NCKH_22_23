import  SelectBasicExample  from '../../components/Shares/search/search__form.js'
import  SelectInput from '../../components/Shares/search/search__input.js'
import './topic.css'


function Topic() {
    return (
      <>
      <div id="main">
        <div id ="topic__name">
          <p>Đề tài nghiên cứu khoa học</p>
          <SelectInput></SelectInput>
        </div>
        <div >
          <SelectBasicExample></SelectBasicExample>
        </div>
      </div>
      </>
    );
  }
  
  export default Topic;