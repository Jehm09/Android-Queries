package com.example.domainqueries.view;

import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.example.domainqueries.R;
import com.example.domainqueries.model.Server;

public class ServersAdapter extends RecyclerView.Adapter<ServersAdapter.ViewHolderData> {

    private Server[] servers;

    public ServersAdapter(Server[] servers) {
        this.servers = servers;
    }

    @NonNull
    @Override
    public ViewHolderData onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(R.layout.servers_recycler, null, false);
        return new ViewHolderData(view);
    }

    @Override
    public void onBindViewHolder(@NonNull ViewHolderData holder, int position) {
        holder.setData(servers[position]);
    }

    @Override
    public int getItemCount() {
        return servers.length;
    }

    public class ViewHolderData extends RecyclerView.ViewHolder {

        private TextView addressTV, sslGradeTV, countryTV, ownerTV;

        public ViewHolderData(@NonNull View itemView) {
            super(itemView);

            addressTV = itemView.findViewById(R.id.historyHostNameTV);
            sslGradeTV = itemView.findViewById(R.id.sslGradeTV);
            countryTV = itemView.findViewById(R.id.countryTV);
            ownerTV = itemView.findViewById(R.id.ownerTV);
        }

        public void setData(Server server) {
            addressTV.setText("Address: " + server.getAddress());
            sslGradeTV.setText("Ssl Grade: " + server.getSsl_grade());
            countryTV.setText("Country: " + server.getCountry());
            ownerTV.setText("Owner: " + server.getOwner());
        }
    }
}
