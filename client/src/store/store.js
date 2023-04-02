import { configureStore } from '@reduxjs/toolkit'
import todoSlice from '../pages/Topic/components/TodoList/todoSlice.js'
import filtersSlice from '../pages/Topic/components/Filters/filtersSlice.js'
import { loginReducer } from '../pages/Login/LoginSlice.js'

const store = configureStore({
  reducer: {
    filters: filtersSlice.reducer,
    todoList: todoSlice.reducer,
    login: loginReducer
  }
})

export default store