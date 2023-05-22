import "./block.css";

const Card = ({ title, children }) => {
  return (
    <div className="card container-fluid">
      <div className="card-header">{title}</div>
      <div style={{padding: "24px"}}>{children}</div>
    </div>
  );
};

export default Card;
