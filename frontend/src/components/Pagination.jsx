export default function Pagination({ page, totalPages, onPageChange }) {
    if (totalPages <= 1) return null

    return (
        <div className="pagination">
            <button
                className="page-btn"
                disabled={page <= 1}
                onClick={() => onPageChange(page - 1)}
            >
                Previous
            </button>
            {Array.from({ length: totalPages }, (_, i) => i + 1).map((p) => (
                <button
                    key={p}
                    className={`page-btn ${p === page ? 'page-active' : ''}`}
                    onClick={() => onPageChange(p)}
                >
                    {p}
                </button>
            ))}
            <button
                className="page-btn"
                disabled={page >= totalPages}
                onClick={() => onPageChange(page + 1)}
            >
                Next
            </button>
        </div>
    )
}