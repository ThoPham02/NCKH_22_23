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
import { CommonCurrentTopicSelector, userSelector } from "../../../store/selectors";
import { useEffect, useState } from "react";
import { fetchCommonDoneTopics, fetchTopics } from "../../../pages/common/Topic/CommonTopicSlice";
import { AdminStatus, LIMIT, StudentLectureStatus } from "../../../const/const";
import { fetchDepartmentCurentTopics } from "../../../pages/department/Topic/DepartmentTopicSlice";

export const SearchWord = Word;
export const SearchStatus = Status;
export const SearchDateTo = DateTo;
export const SearchDateFrom = DateFrom;
export const SearchFaculty = Faculty;
export const SearchDepartment = Department;

export const TopicSearch = (props) => {
    const { pagi, setPagi, isLoading } = props
    const dispatch = useDispatch()
    const [search, setSearch] = useState("")
    const [event, setEvent] = useState(0)
    const [department, setDepartment] = useState(0)
    const [faculty, setFaculty] = useState(0)
    const [status, setStatus] = useState(0)
    const [dateFrom, setDateFrom] = useState("")
    const [dateTo, setDateTo] = useState("")

    const handleSubmitForm = (e) => {
        e.preventDefault()
        setPagi(1)
        dispatch(fetchCommonDoneTopics({
            search: search,
            departmentID: department,
            facultyID: faculty,
            status: status,
            eventID: event,
            timeStart: dateFrom,
            timeEnd: dateTo,
            limit: LIMIT,
            offset: 0,
            isCurrent: 1
        }))
    }
    useEffect(() => {
        dispatch(fetchCommonDoneTopics({
            search: search,
            departmentID: department,
            facultyID: faculty,
            status: status,
            eventID: event,
            timeStart: dateFrom,
            timeEnd: dateTo,
            limit: LIMIT,
            offset: LIMIT * (pagi - 1),
            isCurrent: 1
        }))
        // eslint-disable-next-line
    }, [dispatch, pagi])
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
    const isLoading = useSelector(CommonCurrentTopicSelector).status === 'loading'

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

export const DepartmentCurrentTopicSearch = () => {
    const dispatch = useDispatch()
    const [search, setSearch] = useState("")
    const [status, setStatus] = useState(0)

    const departmentID = useSelector(userSelector).department_id
    const handleSubmitForm = (e) => {
        e.preventDefault()
        dispatch(fetchDepartmentCurentTopics({ departmentID: departmentID, status: status, search: search }))
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

    const handleSubmitForm = (e) => {
        e.preventDefault()
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

export const FacultyCurrentTopicSearch = ({ faculty }) => {
    const dispatch = useDispatch()
    const [search, setSearch] = useState("")
    const [status, setStatus] = useState(0)
    const [department, setDepartment] = useState(0)

    const departmentID = useSelector(userSelector).department_id
    const handleSubmitForm = (e) => {
        e.preventDefault()
        dispatch(fetchDepartmentCurentTopics({ departmentID: departmentID, status: status, search: search }))
    }
    const isLoading = false
    return (
        <Form className="search" onSubmit={handleSubmitForm}>
            <Word search={search} setSearch={setSearch} />
            <Status status={status} setStatus={setStatus} sumStatus={AdminStatus} />
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

export const FacultyDoneTopicSearch = ({ faculty }) => {
    const dispatch = useDispatch()

    const [search, setSearch] = useState("")
    const [event, setEvent] = useState(0)
    const [status, setStatus] = useState(0)
    const [dateFrom, setDateFrom] = useState("")
    const [dateTo, setDateTo] = useState("")
    const [department, setDepartment] = useState(0)

    const handleSubmitForm = (e) => {
        e.preventDefault()
        dispatch()
    }
    const isLoading = false
    return (
        <Form className="search" onSubmit={handleSubmitForm}>
            <Word search={search} setSearch={setSearch} />
            <Event event={event} setEvent={setEvent} />
            <Status status={status} setStatus={setStatus} sumStatus={AdminStatus} />
            <Department
                department={department}
                setDepartment={setDepartment}
                faculty={faculty}
                defaultValue={"Tất cả bộ môn"}
            />
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