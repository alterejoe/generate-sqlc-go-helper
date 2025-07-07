from typing import List, Optional

from sqlalchemy import (
    Boolean,
    DateTime,
    Double,
    ForeignKeyConstraint,
    Index,
    Integer,
    PrimaryKeyConstraint,
    Text,
    text,
)
from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column, relationship
import datetime


class Base(DeclarativeBase):
    pass


class Account(Base):
    __tablename__ = "account"
    __table_args__ = (
        PrimaryKeyConstraint("id", name="account_pkey"),
        {"schema": "budget"},
    )

    id: Mapped[str] = mapped_column(Text, primary_key=True)
    name: Mapped[Optional[str]] = mapped_column(Text)
    balance: Mapped[Optional[float]] = mapped_column(Double(53))
    org_id: Mapped[Optional[str]] = mapped_column(Text)
    org_name: Mapped[Optional[str]] = mapped_column(Text)
    last_updated: Mapped[Optional[datetime.datetime]] = mapped_column(
        DateTime(True), server_default=text("now()")
    )

    transaction: Mapped[List["Transaction"]] = relationship(
        "Transaction", back_populates="account"
    )


class Transaction(Base):
    __tablename__ = "transaction"
    __table_args__ = (
        ForeignKeyConstraint(
            ["account_id"], ["budget.account.id"], name="transaction_account_id_fkey"
        ),
        PrimaryKeyConstraint("id", name="transaction_pkey"),
        Index("transaction_account_id_idx", "account_id"),
        Index("transaction_posted_idx", "transactedat"),
        {"schema": "budget"},
    )

    id: Mapped[str] = mapped_column(Text, primary_key=True)
    account_id: Mapped[str] = mapped_column(Text)
    posted: Mapped[Optional[int]] = mapped_column(Integer)
    amount: Mapped[Optional[str]] = mapped_column(Text)
    description: Mapped[Optional[str]] = mapped_column(Text)
    payee: Mapped[Optional[str]] = mapped_column(Text)
    memo: Mapped[Optional[str]] = mapped_column(Text)
    transactedat: Mapped[Optional[int]] = mapped_column(Integer)
    pending: Mapped[Optional[bool]] = mapped_column(Boolean)
    last_updated: Mapped[Optional[datetime.datetime]] = mapped_column(
        DateTime(True), server_default=text("now()")
    )

    account: Mapped["Account"] = relationship("Account", back_populates="transaction")
