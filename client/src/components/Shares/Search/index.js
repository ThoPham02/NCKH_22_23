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
import { AdminStatus, StudentLectureStatus } from "../../../const/const";
import { fetchDepartmentCurentTopics } from "../../../pages/department/Topic/DepartmentTopicSlice";

export const SearchWord = Word;
export const SearchStatus = Status;
export const SearchDateTo = DateTo;
export const SearchDateFrom = DateFrom;
export const SearchFaculty = Faculty;
export const SearchDepartment = Department;

export const TopicSearch = (props) => {
    const dispatch = useDispatch()

    const [search, setSearch] = useState("")
    const [event, setEvent] = useState(0)
    const [department, setDepartment] = useState(0)
    const [faculty, setFaculty] = useState(0)
    const [status, setStatus] = useState(0)
    const [dateFrom, setDateFrom] = useState("")
    const [dateTo, setDateTo] = useState("")

    const handleSubmitForm = () => {
        dispatch(fetchTopics())
    }
    const isLoading = false
    return (
        <Form className="search" onSubmit={handleSubmitForm}>
            <Word search={search} setSearch={setSearch} />
            <Event event={event} setEvent={setEvent} />
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
            <Status status={status} setStatus={setStatus} sumStatus={AdminStatus} />
            <DateFrom dateFrom={dateFrom} setDateFrom={setDateFrom} />
            <DateTo dateTo={dateTo} setDateTo={setDateTo} />
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
    const [status, setStatus] = useState(0)
    const dispatch = useDispatch()
    const isLoading = useSelector(CommonTopicSelector).status === 'loading'

    const handleSubmitForm = (e) => {
        e.preventDefault()
        dispatch(fetchTopics({
            search: search,
            facultyID: faculty,
            departmentID: department,
            offset: offset,
            limit: limit
        }))
    }

    useEffect(() => {
        dispatch(fetchTopics({}))
    }, [dispatch])

    return (
        <Form className="search" onSubmit={handleSubmitForm}>
            <Word search={search} setSearch={setSearch} />
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
            <Status status={status} setStatus={setStatus} sumStatus={StudentLectureStatus} />
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

export const DepartmentCurrentTopicSearch = (props) => {
    const { departmentID } = props

    const dispatch = useDispatch()

    const [search, setSearch] = useState("")
    const [status, setStatus] = useState(0)

    const handleSubmitForm = () => {
        dispatch(fetchTopics(fetchDepartmentCurentTopics({departmentID: departmentID, status: status, search: search})))
    }
    const isLoading = false
    return (
        <Form className="search" onSubmit={handleSubmitForm}>
            <Word search={search} setSearch={setSearch} />
            <Status status={status} setStatus={setStatus} sumStatus={AdminStatus} />
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

export const DepartmentDoneTopicSearch = (props) => {
    const dispatch = useDispatch()

    const [search, setSearch] = useState("")
    const [event, setEvent] = useState(0)
    const [status, setStatus] = useState(0)
    const [dateFrom, setDateFrom] = useState("")
    const [dateTo, setDateTo] = useState("")

    const handleSubmitForm = () => {
        dispatch()
    }
    const isLoading = false
    return (
        <Form className="search" onSubmit={handleSubmitForm}>
            <Word search={search} setSearch={setSearch} />
            <Event event={event} setEvent={setEvent} />
            <Status status={status} setStatus={setStatus} sumStatus={AdminStatus} />
            <DateFrom dateFrom={dateFrom} setDateFrom={setDateFrom} />
            <DateTo dateTo={dateTo} setDateTo={setDateTo} />
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