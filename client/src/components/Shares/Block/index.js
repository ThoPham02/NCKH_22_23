import "./block.css"

const Block = ({title, children}) => {
    return (
        <div className="block container-fluid">
            <div className="block-header">
                {title}
            </div>
            {children}
        </div>
    )
}

export default Block;