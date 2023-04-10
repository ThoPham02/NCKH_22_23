import { configureStore } from '@reduxjs/toolkit'
import { loginReducer } from '../pages/Login/LoginSlice.js'
import { FilterReducer } from '../components/Shares/Search/SearchSlice.js'
import { DepartmentReducer } from '../components/Shares/Search/SearchDepartment/DepartmentSlice.js'
import { FaculityReducer } from '../components/Shares/Search/SearchFaculity/FaculitySlice.js'

const store = configureStore({
  reducer: {
    faculity: FaculityReducer,
    department: DepartmentReducer,
    topicFilter: FilterReducer,
    login: loginReducer
  }
})

export default store