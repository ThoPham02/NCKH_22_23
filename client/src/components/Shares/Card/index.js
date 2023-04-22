import "./block.css"

const Card = ({title, children}) => {
    return (
        <div className="card container-fluid">
            <div className="card-header">
                {title}
            </div>
            {children}
        </div>
    )
}

export default Card;