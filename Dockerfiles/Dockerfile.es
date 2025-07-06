FROM opensearchproject/opensearch:3


# ローカル環境用のためセキュリティプラグインをアンインストール
RUN /usr/share/opensearch/bin/opensearch-plugin remove opensearch-security
# プラグインのインストール
RUN /usr/share/opensearch/bin/opensearch-plugin install analysis-kuromoji
RUN /usr/share/opensearch/bin/opensearch-plugin install analysis-icu
