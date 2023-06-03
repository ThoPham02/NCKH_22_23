import { Button, Form } from "react-bootstrap";

import DateFrom from "./DateFrom";
import DateTo from "./DateTo";
import Department from "./Department";
import Faculty from "./Faculty";
import Status from "./Status";
import Word from "./Word";
import Loading from "../Loading";
import Event from "./Event";
import { useDispatch, useSelector } from "react-redux";
import { CommonTopicSelector } from "../../../store/selectors";
import { useEffect, useState } from "react";
import { fetchTopics } from "../../../pages/common/Topic/CommonTopicSlice";

export const SearchWord = Word;
export const SearchStatus = Status;
export const SearchDateTo = DateTo;
export const SearchDateFrom = DateFrom;
export const SearchFaculty = Faculty;
export const SearchDepartment = Department;

export const TopicSearch = (props) => {
    const { handleSubmitForm, searchRef, departmentRef, statusRef, dateFromRef, dateToRef, eventRef, isLoading, faculty, setFaculty } = props

    return (
        <Form className="search" onSubmit={handleSubmitForm}>
            <Word searchRef={searchRef} />
            <Event eventRef={eventRef} />
            <Faculty
                faculty={faculty}
                setFaculty={setFaculty}
            />
            <Department
                departmentRef={departmentRef}
                faculty={faculty}
                defaultValue={"Tất cả bộ môn"}
            />
            <Status statusRef={statusRef} />
            <DateFrom dateFromRef={dateFromRef} />
            <DateTo dateToRef={dateToRef} />
            <Form.Group>
                <Button
                    variant="primary"
                    type="submit"
                    className="search-submit"
                >
                    {isLoading ? <Loading></Loading> : <></>}
                    Tìm kiếm
                </Button>
            </Form.Group>
        </Form>
    )

}

export const CurrentTopicsSearch = (props) => {
    const { offset, limit } = props

    const [faculty, setFaculty] = useState(0)
    const [search, setSearch] = useState("")
    const [department, setDepartment] = useState(0)
    const dispatch = useDispatch()
    const isLoading = useSelector(CommonTopicSelector).status === 'loading'

    const handleSubmitForm = (e) => {
        e.preventDefault()
        dispatch(fetchTopics({search: search, facultyID: faculty, departmentID: department, offset: offset, limit: limit}))
    }

    useEffect(() => {
        dispatch(fetchTopics({}))
    }, [dispatch])

    return (
        <Form className="search" onSubmit={handleSubmitForm}>
            <Word search={search} setSearch={setSearch}/>
            <Faculty
                faculty={faculty}
                setFaculty={setFaculty}
            />
            <Department
                department={department}
                setDepartment={setDepartment}
                faculty={faculty}
                defaultValue={"Tất cả bộ môn"}
            />
            <Form.Group>
                <Button
                    variant="primary"
                    type="submit"
                    className="search-submit"
                >
                    {isLoading ? <Loading></Loading> : <></>}
                    Tìm kiếm
                </Button>
            </Form.Group>
        </Form>
    )
}