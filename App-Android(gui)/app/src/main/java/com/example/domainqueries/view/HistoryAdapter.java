package com.example.domainqueries.view;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.example.domainqueries.R;
import com.example.domainqueries.model.Server;

public class HistoryAdapter extends RecyclerView.Adapter<HistoryAdapter.ViewHolderData> {

    private String []items;

    public HistoryAdapter(String[] items) {
        this.items = items;
    }

    @NonNull
    @Override
    public ViewHolderData onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.history_recycler, null, false);
        return new ViewHolderData(view);
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolderData holder, int position) {
        holder.setData(items[position]);
    }

    @Override
    public int getItemCount() {
        return items.length;
    }

    public class ViewHolderData extends RecyclerView.ViewHolder {

        private TextView historyHostNameTV;

        public ViewHolderData(@NonNull View itemView) {
            super(itemView);

            historyHostNameTV = itemView.findViewById(R.id.historyHostNameTV);
        }

        public void setData(String item) {
            historyHostNameTV.setText(item);
        }
    }
}
