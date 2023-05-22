import "./style.css"

const SubCard = ({title, children, style}) => {
    return (
        <div className="subcard container-fluid">
            <div className="subcard-header" style={style}>
                {title}
            </div>
            {children}
        </div>
    )
}

export default SubCard;