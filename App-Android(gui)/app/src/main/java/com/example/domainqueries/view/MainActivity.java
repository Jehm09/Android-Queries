package com.example.domainqueries.view;

import androidx.appcompat.app.AppCompatActivity;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import android.os.Bundle;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.TextView;

import com.example.domainqueries.R;
import com.example.domainqueries.control.MainController;

public class MainActivity extends AppCompatActivity {

    private TextView titleTV, sslGradeTV, previousSslTV, isDownTV, serverChangedTV;
    private ImageView logoIV;
    private RecyclerView serversRecycler;
    private RecyclerView historyRecycler;
    private EditText hostNamePT;
    private Button searchBtn;

    private MainController mainController;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        titleTV = findViewById(R.id.titleTV);
        sslGradeTV = findViewById(R.id.sslGradeTV);
        previousSslTV = findViewById(R.id.previousSslTV);
        isDownTV = findViewById(R.id.isDownTV);
        serverChangedTV = findViewById(R.id.serverChangedTV);
        hostNamePT = findViewById(R.id.hostNamePT);
        searchBtn = findViewById(R.id.searchBtn);

        logoIV = findViewById(R.id.logoIV);
        serversRecycler = findViewById(R.id.serversRecycler);
        serversRecycler.setLayoutManager(new LinearLayoutManager(this, LinearLayoutManager.VERTICAL,false));

        historyRecycler = findViewById(R.id.historyRecycler);
        historyRecycler.setLayoutManager(new LinearLayoutManager(this, LinearLayoutManager.VERTICAL,false));

        mainController = new MainController(this);
    }

    public RecyclerView getHistoryRecycler() {
        return historyRecycler;
    }

    public void setHistoryRecycler(RecyclerView historyRecycler) {
        this.historyRecycler = historyRecycler;
    }

    public EditText getHostNamePT() {
        return hostNamePT;
    }

    public void setHostNamePT(EditText hostNamePT) {
        this.hostNamePT = hostNamePT;
    }

    public Button getSearchBtn() {
        return searchBtn;
    }

    public void setSearchBtn(Button searchBtn) {
        this.searchBtn = searchBtn;
    }

    public MainController getMainController() {
        return mainController;
    }

    public void setMainController(MainController mainController) {
        this.mainController = mainController;
    }

    public TextView getTitleTV() {
        return titleTV;
    }

    public void setTitleTV(TextView titleTV) {
        this.titleTV = titleTV;
    }

    public TextView getSslGradeTV() {
        return sslGradeTV;
    }

    public void setSslGradeTV(TextView sslGradeTV) {
        this.sslGradeTV = sslGradeTV;
    }

    public TextView getPreviousSslTV() {
        return previousSslTV;
    }

    public void setPreviousSslTV(TextView previousSslTV) {
        this.previousSslTV = previousSslTV;
    }

    public TextView getIsDownTV() {
        return isDownTV;
    }

    public void setIsDownTV(TextView isDownTV) {
        this.isDownTV = isDownTV;
    }

    public TextView getServerChangedTV() {
        return serverChangedTV;
    }

    public void setServerChangedTV(TextView serverChangedTV) {
        this.serverChangedTV = serverChangedTV;
    }

    public ImageView getLogoIV() {
        return logoIV;
    }

    public void setLogoIV(ImageView logoIV) {
        this.logoIV = logoIV;
    }

    public RecyclerView getServersRecycler() {
        return serversRecycler;
    }

    public void setServersRecycler(RecyclerView serversRecycler) {
        this.serversRecycler = serversRecycler;
    }
}
