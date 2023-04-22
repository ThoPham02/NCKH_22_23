import { configureStore } from '@reduxjs/toolkit'
import { LoginReducer } from '../pages/common/Login/loginSlice'
import { DepartmentReducer } from '../components/Shares/Search/Department/DepartmentSlice'
import { FaculityReducer } from '../components/Shares/Search/Faculty/FaculitySlice'

const store = configureStore({
  reducer: {
    login: LoginReducer,
    department: DepartmentReducer,
    faculty: FaculityReducer,
  }
})

export default store