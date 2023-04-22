import "./style.css"

const SubCard = ({title, children}) => {
    return (
        <div className="subcard container-fluid">
            <div className="subcard-header">
                {title}
            </div>
            {children}
        </div>
    )
}

export default SubCard;