import { useSelector } from "react-redux"
import { departmentSelector, facultySelector } from "../store/selectors"

export const useDepartmentFaculty = (departmentID) => {
    const listDepartment = useSelector(departmentSelector)
    const listFacul = useSelector(facultySelector)

    const department = listDepartment.find(item => item.id === departmentID)
    const faculty = listFacul.find(item => item.id === department.facultyID)

    return {department: department.name, faculty: faculty.name}
}