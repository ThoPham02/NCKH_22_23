const TableHeader = ({listKey}) => {
  return (
    <thead>
      <tr>
        {listKey.map((key, index) => {
            return (<th style={{alignItems:"center"}}key={index}>{key}</th>)
        })}
      </tr>
    </thead>
  );
};

export default TableHeader;
