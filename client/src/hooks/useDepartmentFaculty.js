import { useDispatch, useSelector } from "react-redux"
import { departmentSelector, facultySelector } from "../store/selectors"
import { fetchFaculty } from "../components/Shares/Search/Faculty/FacultySlice";
import { fetchDepartment } from "../components/Shares/Search/Department/DepartmentSlice";
import { useEffect } from "react";

export const useDepartmentFaculty = (departmentID) => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(fetchFaculty())
        dispatch(fetchDepartment())
    }, [dispatch]);
    const listDepartment = useSelector(departmentSelector)
    const listFacul = useSelector(facultySelector)

    if (listDepartment.length > 0 && listFacul.length > 0) {
        const department = listDepartment.find(item => item.id === departmentID)
        const faculty = listFacul.find(item => item.id === department.facultyID)
        return { department: department.name, faculty: faculty.name }
    }
    return {}
}