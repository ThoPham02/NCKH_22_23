import "./adminlayout.css"

const AdminLayout = ({ children }) => {
    return (
        <div className="admin">
            {children}
        </div>
    );
};

export default AdminLayout;
