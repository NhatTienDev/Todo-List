export default function ConfirmModal({ title, onConfirm, onCancel }) {
    return (
        <div className="modal-overlay">
            <div className="modal-content" onClick={(e) => e.stopPropagation()}>
                <p className="modal-message">{title}</p>
                <div className="modal-actions">
                    <button className="btn-danger" onClick={onConfirm}>
                        Delete
                    </button>
                    <button className="btn-cancel" onClick={onCancel}>
                        Cancel
                    </button>
                </div>
            </div>
        </div>
    );
}