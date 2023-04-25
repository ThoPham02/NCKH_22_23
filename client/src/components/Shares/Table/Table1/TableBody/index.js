const TableBody = ({ listKey, listItem }) => {
  return (
    <tbody>
      {listItem.map((item, index) => (
        <tr key={index}>
          <td>{index + 1}</td>
          {listKey.map((key, index) => (
            <td key={index}>{item[key]}</td>
          ))}
          
        </tr>
      ))}
    </tbody>
  );

};

export default TableBody;
