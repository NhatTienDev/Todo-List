export default function SearchBar({
    search,
    onSearchChange,
    filterStatus,
    onFilterChange,
}) {
    return (
        <div className="search-bar">
            <input
                type="text"
                placeholder="Search by title"
                value={search}
                onChange={(e) => onSearchChange(e.target.value)}
            />
            <select
                value={filterStatus}
                onChange={(e) => onFilterChange(e.target.value)}
            >
                <option value="">All</option>
                <option value="false">Pending</option>
                <option value="true">Completed</option>
            </select>
        </div>
    );
}