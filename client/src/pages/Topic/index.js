import SelectBasicExample from '../../components/Shares/search/search__form.js'
import SelectInput from '../Topic/components/Filters/index.js'
import AddTodo from "../Topic/components/TodoList/index.js"
import '../Topic/components/TodoList'
import './topic.css'
import { useSelector } from 'react-redux';
import { todoRemainingSelector } from '../../store/selectors.js'
import TodoTable from '../Topic/components/Todo/index.js'

function Topic() {
  const todoList = useSelector(todoRemainingSelector)
  var openSearch = false
  return (
    <>
      <div id="topic__name">
        <p>Đề tài nghiên cứu khoa học</p>
      </div>
       <div>
        <AddTodo></AddTodo>
      </div> 
      <div id="topic__search">
        <div id="topic__search-1">
          <SelectBasicExample openSearch={openSearch}></SelectBasicExample>
        </div>
        <div id="topic__search-2">
          <SelectInput></SelectInput>
        </div>

        <div id="topic__table">
          <TodoTable todoList={todoList}></TodoTable>
        </div>
      </div>

    </>
  );
}

export default Topic;